
> 指纹是Json格式,如下指纹案例

```json
"Shiro": {
      "cookies": {
        "rememberMe": ""
      }
}
```

```json
 "PHPMyAdmin": {
      "html": "<title>phpMyAdmin </title>",
      "html": "/themes/pmahomme/img/logo_right.png"
} 
```

```json
 "Kibana": {
      "headers": {
        "kbn-name": "kibana",
        "kbn-version": "^([\\d.]+)$\\;version:\\1"
      },
      "html": "<title>Kibana</title>"
    }
```


| 参数                  | 说明                                                         | 
| :-------------------- | :----------------------------------------------------------- |
| html                  | 正则匹配网页的html                                    |
| headers                  | 正则匹配网页的headers                                   |
| cookies                  | 正则匹配网页的cookies                                    |
