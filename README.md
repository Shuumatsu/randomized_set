#randomizedSet

A data structure that supports insert, delete, search and getRandom in constant time.

#USAGE

```
func New() *Set 
```
```
func (set *Set) Add(element interface{}) 
```
```
func (set *Set) Search(element interface{}) bool 
```
```
func (set *Set) GetRandom() interface{} 
```
```
func (set *Set) Delete(element interface{}) 
```

#LICENSE
The MIT License (MIT)

Inspired by [Ronny.Pena][1] and [1337c0d3r][2]


  [1]: https://discuss.leetcode.com/user/ronny-pena
  [2]: https://discuss.leetcode.com/user/1337c0d3r
