package helper

import (
    "hash/fnv"
)

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func MapKeysInSlice(dict map[string]interface{}, list []string) bool {
    allMapKeysInArray := false;
    for k, _ := range dict {
        allMapKeysInArray = false;
        for _, v := range list {
            if v == k {
                allMapKeysInArray = true;
            }
        }
        if !allMapKeysInArray {
            allMapKeysInArray = false;
            break;
        }


    }
    return allMapKeysInArray
}


func HashString(s string) int {
    h := fnv.New32a()
    h.Write([]byte(s))
    return int(h.Sum32())
}

func GetIndex(fielddata interface{}) int {
    var val int;

    switch res := fielddata.(type) {
    case int:
        val = res
    case string:
        val = HashString(res)	
    }
    return val
}