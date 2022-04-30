package biz

const menu = `
{
  "button": [
    {
      "type": "view",
      "name": "百度一下",
      "url": "https://www.baidu.com/"
    },
    {
      "type": "view",
      "name": "开发文档",
      "url": "https://developers.weixin.qq.com/doc"
    },
    {
      "name": "菜单",
      "sub_button": [
        {
          "type": "view",
          "name": "搜索",
          "url": "https://www.baidu.com/"
        },
        {
          "type": "view",
          "name": "wxa",
          "url": "http://mp.weixin.qq.com"
        },
        {
          "type": "click",
          "name": "赞一下我们",
          "key": "V1001_GOOD"
        }
      ]
    }
  ]
}
`
