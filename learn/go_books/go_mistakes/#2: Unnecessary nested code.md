# Unnecessary nested code（避免代码嵌套）
1. 嵌套太多，没法看清逻辑
```
func join(s1, s2 string, max int) (string, error) {
    // 一些额外情况的判断与处理
    if s1==""{
        return "", errors.New("s1 is empty")
    }

    if s2==""{
        return "", errors.New("s2 is empty")
    }

    concat, err := concatenate(s1, s2)
    if err != nil {
        return "", err
    }

    if len(concat) > max {
        return concat[:max], nil
    }

    return concat, nil
}
```