<div align="center">

# VanVo - Ngôn ngữ lập trình với cú pháp thuần Việt
[![](https://raw.githubusercontent.com/webuild-community/badge/master/svg/made.svg)](https://vi.wikipedia.org/wiki/Vi%E1%BB%87t_Nam)
[![](https://img.shields.io/badge/made_with-Go-blue?logo=go)](https://go.dev/)
[![bun-dau](https://img.shields.io/badge/b%C3%BAn%20%C4%91%E1%BA%ADu-m%E1%BA%AFm%20t%C3%B4m-red?logo=undertale)](https://vi.wikipedia.org/wiki/B%C3%BAn_%C4%91%E1%BA%ADu_m%E1%BA%AFm_t%C3%B4m)

</div>


VanVo (Văn Vở) là ngôn ngữ lập trình được thiết kế với cú pháp thuần Việt, nhưng không chỉ đơn giản là dịch lại một cách gượng gạo từ các ngôn ngữ khác, mình cố gắng để thiết kế một ngôn ngữ sao cho các câu lệnh trông tự nhiên như tiếng Việt nhất có thể, thuận tiện cho người Việt nhất có thể. Điển hình là bạn có thể đặt tên định danh có khoảng trắng như `số nguyên tố`.

Ngôn ngữ sẽ có hơi thiên hướng toán học, những bạn học toán có thể sẽ cảm thấy quen thuộc hơn, vì mình định hướng ngôn ngữ được sử dụng như một CAS (Computer Algebra System). Tuy nhiên VanVo vẫn có thể dùng như một ngôn ngữ đa mục đích (General-purpose language) thông thường.

## Điểm qua một số tính năng của VanVo

-   Hỗ trợ những câu lệnh rẽ nhánh, cấu trúc lặp, cấu trúc dữ liệu và phép toán cơ bản.
-   Hỗ trợ phân số và số phức.
-   Có thể đặt tên định danh có khoảng trắng như `số nguyên tố`.
-   Không cần `;` ở cuối mỗi câu lệnh, và các khối lệnh sẽ được xác định bởi mức thụt dòng (indent level) như Python.
-   Phép nhân giữa hằng số, biến và mở ngoặc có thể lược bỏ, ví dụ `2x(x-1)` sẽ tương đương với `2*x*(x-1)`.
-   List comprehension như `{ n*m | n thuộc [1..10], m thuộc [1..10], n != m }`
-   Lazy evaluation.
-   Các thao tác và phép toán trên tập hợp như hội, giao, hiệu, tích Descartes.
-   Gạch chân chính xác vị trí có lỗi khi chạy chương trình.

## Cài đặt

Nếu bạn đã tải [Go](https://go.dev/) thì cách đơn giản nhất để cài đặt là clone và build trực tiếp từ source như sau

```bash
git clone https://github.com/balldk/vanvo
cd vanvo
go install .
```

Ngoài ra bạn có thể tải file thực thi tại đây [Releases v0.1.0](https://github.com/balldk/vanvo/releases/tag/v0.1.0).

**Cài đặt language support:** Để ngôn ngữ có syntax highlight và snippets trên VSCode thì bạn hãy qua đây cài thêm extension nhé: [github.com/balldk/vscode-vanvo](https://github.com/balldk/vscode-vanvo).

Để chạy chương trình ta dùng lệnh

```vanvo
vanvo program.vv
```

Trong đó `program.vv` là file chứa nội dung code. Nói chung file extension của VanVo là `.vv`.

## Một số ví dụ minh họa

**Ví dụ 1:** Xét tính chia hết của n cho 2 và 3, với n là các số nguyên trong khoảng $[1,100]$

```vanvo
cho A = [1..100]

với mỗi n thuộc A:
	nếu n % 2 == 0:
		xuất n, "chia hết cho 2"
	còn nếu n % 3 == 0:
		xuất n, "chia hết cho 3"
	còn không:
		xuất n, "không chia hết cho cả 2 và 3"
```

**Ví dụ 2:** Tính giá trị của hàm hợp, với $(f.g)(x) = f(g(x))$

```vanvo
cho f(x) = 2x(x^2 - 2x)(3x - 5)
cho g(x) = 5x

cho a = 5
xuất f.g(a)
```

`2x(x^2 - 2x)(3x - 5)` là cách viết ngắn gọn hơn của `2*x*(x^2 - 2*x)*(3*x - 5)`.

**Ví dụ 3:** Sử dụng list comprehension để định nghĩa mảng vô hạn các phần tử, với `fib[i]` là phần tử thứ `i` trong dãy fibonacci

```vanvo
cho fib = {0, 1, 1} + { fib[n-1] + fib[n-2] | n thuộc [3..] }

với mỗi n thuộc fib:
    xuất n
```

List comprehension có tính "lazy", tức là chỉ khi ta cần dùng phần tử nào trong mảng thì phần tử đó mới được tính ra, do đó ta có thể dễ dàng định nghĩa ra mảng vô hạn phần tử trong VanVo.

**Ví dụ 4:** Hệ thống báo lỗi của chương trình

![screenshot_error](/examples/screenshot_error.png)

## Tài liệu

Tài liệu được viết chi tiết ở đây: [balldk.github.io/posts/vanvo/](https://balldk.github.io/posts/vanvo/)
