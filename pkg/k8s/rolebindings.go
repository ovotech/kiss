package k8s

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/ovotech/kiss/pkg/authz/utils"
	"github.com/rs/zerolog/log"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var client *kubernetes.Clientset = nil

func isSupportedRolebinding(name string, prefix string) bool {
	return strings.HasPrefix(name, prefix)
}

func getClientset(configPath string) *kubernetes.Clientset {

	if client != nil {
		return client
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Error().Msg("failed to load kubeconfig")
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	client = clientset
	return clientset

}

func getRoleBinding(clientset *kubernetes.Clientset, namespace string, roleBindingPrefix string) (*v1.RoleBinding, error) {
	roleBindingName := roleBindingPrefix + namespace
	return clientset.RbacV1().RoleBindings(namespace).Get(context.TODO(), roleBindingName, metav1.GetOptions{})
}
func VerifyNamespaceClaims(namespace string, claims []string, roleBindingPrefix string, namespaceRegex string, kubeconfigPath string) (bool, error) {
	clientset := getClientset(kubeconfigPath)
	roleBinding, err := getRoleBinding(clientset, namespace, roleBindingPrefix)
	if err != nil {
		// Could not get roleBinding
		return false, fmt.Errorf("failed to get rolebinding name %s from k8s. error: %w", roleBindingPrefix, err)
	}
	// Check if rolebinding has required prefix
	if isSupportedRolebinding(roleBinding.Name, roleBindingPrefix) {

		// compare claims to Rolebinding subjects
		// presort to improve time complexity
		sort.Strings(claims)
		if err != nil {
			return false, err
		}
		for _, subject := range roleBinding.Subjects {
			for _, claim := range claims {

				// Remove prefix from subject name
				fullName := strings.Split(subject.Name, roleBindingPrefix)
				nameWithPrefix := fullName[len(fullName)-1]
				var name string
				var err error
				if namespaceRegex != "" {

					name, err = utils.ExtractNamespaceFromClaim(namespaceRegex, nameWithPrefix)
					if err != nil {
						log.Debug().Msgf("failed to extract rolebinding name %s using regex %s.", nameWithPrefix, namespaceRegex)
						return false, err
					}
				} else {
					name = nameWithPrefix
					log.Debug().Msgf("No namespace regex specified. Skipping regex extraction.")
				}

				if name == claim {
					return true, nil
				}
			}
		}
	}

	return false, err
}
