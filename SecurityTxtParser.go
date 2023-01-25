package SecurityTxtParser 

import (
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

type KV struct {
    Key   string
    Value string
}

type SecurityTxt struct {
    Contact []string
    Expires []string
    Encryption []string
    Acknowledgments []string
    PreferredLanguages []string
    Canonical []string
    Policy []string
    Hiring []string
}

func ParseTxt(txt string) (SecurityTxt, error) {
    kvs := getLinesWithKeys(txt)
    st := transformToSecurityTxt(kvs)
    return st, nil
}

func transformToSecurityTxt(kvs []KV) SecurityTxt {
    preferredLanguages := splitEachRemoveDuplicates(
        filterForKey(kvs, "Preferred-Languages"), ", ",
    )

    st := SecurityTxt{
        Contact: filterForKey(kvs, "Contact"),
        Expires: filterForKey(kvs, "Expires"),
        Encryption: filterForKey(kvs, "Encryption"),
        Acknowledgments: filterForKey(kvs, "Acknowledgments"),
        PreferredLanguages: preferredLanguages,
        Canonical: filterForKey(kvs, "Canonical"),
        Policy: filterForKey(kvs, "Policy"),
        Hiring: filterForKey(kvs, "Hiring"),
    }
    return st
}

func getLinesWithKeys(text string) []KV {
    pattern := regexp.MustCompile("(.*): (.*)")
    key_lines := pattern.FindAllStringSubmatch(text, -1)

    var key_value []KV
    for _, element := range key_lines {
        
        key_value = append(key_value, KV{Key: element[1], Value: element[2]}) 
    }
    return key_value
}

// Splits the strings in the string array according to separator, then removes the duplicates
func splitEachRemoveDuplicates(s_arr []string, sep string) []string {
    var out []string
    for _, item := range s_arr {
        parsedItems := strings.Split(item, sep)

        for _, currentItem := range parsedItems {
            if !slices.Contains(out, currentItem) {
                out = append(out, currentItem)
            }
        }
    }
    return out
}

func filterForKey(kvs []KV, key string) []string {
    var out []string
    for _, kv := range kvs {
        if kv.Key == key {
            out = append(out, kv.Value)
        }
    }
    return out
}
