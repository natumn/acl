
# ACL 5/23 3週目

koki natsume

---

### 今日やること

- parserの実装を通してhaskellの文法を学ぶ
- 今回はhaskellの文法を学ぶことに重きをおく |

---

### 参考

[school-of-haskell](https://www.schoolofhaskell.com/school/starting-with-haskell/basics-of-haskell/8_Parser)

---

[なぜHaskellを学ぶのか](https://qiita.com/arowM/items/0305d4f439752f285438)

--- 

### Parse Tree

```haskell:ParseTree.hs
data Tree = SumNode Operator Tree Tree
          | ProdNode Operator Tree Tree
          | AssignNode String Tree
          | UnaryNode Operator Tree
          | NumNode Double
          | VarNode String
  deriving Show
```

+++

### 直積型

内部に値を持つ型のこと。他言語の構造体みたいな。

```haskell:Ex.hs
data 型 = コンストラクタ [フィールドの型 ...]
data Point = Point Int Int
```

+++

### 直和型

列挙型にフィールドを付加することで、複数の直積型を定義したもの。

```haskell:Ex.hs
data 型 = コンストラクタ [フィールドの型 ...] | コンストラクタ [フィールドの型 ...] [| ...]
data Foo = Bar Int Int | Baz Int Int Int
```

+++

### 型クラスとは

共通のメソッドを提供する型の集合のこと。

+++
### 例

#### Eqクラス
```haskell:Eq.hs
(==) :: a -> a -> Bool
(/=) :: a -> a -> Bool
```

#### Showクラス
```haskell:Show.hs
show :: a -> String
```

+++

### 関数定義
```haskell:FuncDef.hs
max :: Num a => [a] -> a
max = 処理を書く
```

---


### tokenize

```haskell:Tokenize.hs
tokenize :: String -> [Token]
tokenize [] = []
tokenize (c : cs)
    | elem c "+-*/" = TokOp (operator c) : tokenize cs
    | c == '='  = TokAssign : tokenize cs
    | c == '('  = TokLParen : tokenize cs
    | c == ')'  = TokRParen : tokenize cs
    | isDigit c = number c cs
    | isAlpha c = identifier c cs
    | isSpace c = tokenize cs
| otherwise = error $ "Cannot tokenize " ++ [c]
```

+++

### パターンマッチ

``` haskell:PM.hs
tokenize [] = []
tokenize (c : cs) = 
```

+++ 

### ガード等式

``` haskell:Gad.hs
tokenize (c : cs)
    | elem c "+-*/" = TokOp (operator c) : tokenize cs
    | c == '='  = TokAssign : tokenize cs
    | c == '('  = TokLParen : tokenize cs
    | c == ')'  = TokRParen : tokenize cs
    | isDigit c = number c cs
    | isAlpha c = identifier c cs
    | isSpace c = tokenize cs
| otherwise = error $ "Cannot tokenize " ++ [c]
```

+++

### リストあれこれ

#### ++
++ によりリスト同士を結合
#### : 
: によりリストの先頭に要素を挿入

---

### expression
```haskell:Expression.hs

expression :: [Token] -> (Tree, [Token])
expression toks =
   let (termTree, toks') = term toks
   in
      case lookAhead toks' of
         (TokOp op) | elem op [Plus, Minus] ->
            let (exTree, toks'') = expression (accept toks')
            in (SumNode op termTree exTree, toks'')
         TokAssign ->
            case termTree of
               VarNode str ->
                  let (exTree, toks'') = expression (accept toks')
                  in (AssignNode str exTree, toks'')
               _ -> error "Only variables can be assigned to"
         _ -> (termTree, toks')
```

+++

### let
変数などを定義してから使用。逆にwhereがある
doあり
```
main = do
    let a = 1
        b = 2
        c = a + b
    print c
```

doがなければ最後にinが必要。
doなし

```
main =
    let a = 1
        b = 2
        c = a + b in
    print c
```

+++ 

### case of

関数の中でパターンマッチを行う。=ではなく->を使う。
```haskell:Fact.hs
fact n = case n of
    0 -> 1
    _ | n > 0 -> n * fact (n - 1)

main = do
    print $ fact 5
```

---


## これでコードが全部読めるはず

---

###  ちなみに実行すると

```haskell:Exec.hs
stack exec hs-parser-exe
AssignNode "x1" (ProdNode Div (UnaryNode Minus (NumNode 15.0)) (SumNode Plus (NumNode 2.0) (VarNode "x2")))
```

---

## ご静聴ありがとうござました！
