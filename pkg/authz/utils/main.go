package utils

import (
	"fmt"
	"regexp"
)

func ExtractNamespaceFromClaim(namespaceRegex string, rawNamespace string) (string, error) {
	re := regexp.MustCompile(namespaceRegex)
	matches := re.FindStringSubmatch(rawNamespace)
	if len(matches) != 2 {
		return "", fmt.Errorf("failed to extract namespace from '%s' using regexp `%s`",
			rawNamespace,
			namespaceRegex,
		)

	}
	return matches[1], nil
}

// Extracts namespaces from raw claims array using given regex
func ExtractNamespacesFromClaims(namespacesRegex string, rawNamespaces []string) ([]string, error) {

	re := regexp.MustCompile(namespacesRegex)
	var namespaces []string
	for _, n := range rawNamespaces {
		matches := re.FindStringSubmatch(n)
		if len(matches) != 2 {
			return nil, fmt.Errorf("failed to extract namespace from '%s' using regexp `%s`",
				n,
				namespacesRegex,
			)

		}
		namespaces = append(namespaces, matches[1])
	}
	return namespaces, nil

}
