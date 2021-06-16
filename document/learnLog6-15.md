js中文正则匹配

```js
var pattern = /[小][班]/;
var str = "小提琴班练习课";
console.log(pattern.test(str));
// false

var pattern = /[小][班]/;
var str = "小提琴小班练习课";
console.log(pattern.test(str));
// true
```