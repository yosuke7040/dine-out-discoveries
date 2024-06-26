Golangのcollyパッケージでスクレイピングを行う際、セレクタの使い分けは重要です。以下に、各セレクタの意味と使い方について説明します。

1. `#` (IDセレクタ):
   - HTML要素の`id`属性を指定する場合に使用します。
   - 例: `"#rdnavi-review"`は、`id`属性が`"rdnavi-review"`である要素を選択します。

2. `.` (クラスセレクタ):
   - HTML要素の`class`属性を指定する場合に使用します。
   - 例: `".rstdtl-navi__total-count"`は、`class`属性が`"rstdtl-navi__total-count"`である要素を選択します。

3. `>` (子孫セレクタ):
   - 親要素の直接の子要素を選択する場合に使用します。
   - 例: `"span.rstdtl-navi__total-count > em"`は、`class`属性が`"rstdtl-navi__total-count"`である`<span>`要素の直接の子要素である`<em>`要素を選択します。

4. ` ` (子孫セレクタ):
   - 親要素の子孫要素を選択する場合に使用します。
   - 例: `"div.rstdtl-rvwlst div.js-rvw-item-clickable-area"`は、`class`属性が`"rstdtl-rvwlst"`である`<div>`要素の子孫要素で、`class`属性が`"js-rvw-item-clickable-area"`である`<div>`要素を選択します。

5. `element1, element2` (複数セレクタ):
   - 複数の要素を選択する場合に使用します。
   - 例: `"h1, h2, h3"`は、`<h1>`、`<h2>`、`<h3>`要素をすべて選択します。

6. `[attribute]` (属性セレクタ):
   - 特定の属性を持つ要素を選択する場合に使用します。
   - 例: `"a[href]"`は、`href`属性を持つ`<a>`要素を選択します。

これらのセレクタを組み合わせることで、目的の要素を正確に選択することができます。スクレイピングする際は、対象のWebサイトのHTML構造を確認し、適切なセレクタを使用することが重要です。

また、セレクタの記述方法は、collyパッケージだけでなく、他のスクレイピングライブラリやCSSセレクタの一般的な規則に従っています。これらの規則を理解することで、より効果的にスクレイピングを行うことができます。
