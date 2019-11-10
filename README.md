|   关键字   |  英文描述    |   中文描述   |
| ---- | ---- | ---- |
|    s  |   real, single precision   |    实数单精度  |
|   d   |     eal, double precision |    实数双精度  |
|    c  |    complex, single precision |  复数单精度    |
|    z  |    complex, double precision |  复数双精度    |
|     ge |  general matrix	  | 一般矩阵     |
|   gb   |  general band matrix  |    一般带状矩阵  |
|    sy	  |   symmetric matrix |   对称矩阵   |
|   sp   |   symmetric matrix (packed storage) |   对称矩阵(压缩存储)   |
|    sb  |  symmetric band matrix |   对称带状矩阵  |
|  mv    |  matrix-vector product  |   矩阵向量乘积   |
|   sv	   |  solving a system of linear equations with a single unknown vector  |    求解含有一个未知向量解线性方程组  |
|   mm   |  matrix-matrix product	  | 矩阵矩阵乘积     |
|   sm	   |solving a system of linear equations with multiple unknown vectors    |    求解含有多个未知向量的线性方程组  |
|   	tr   | triangular matrix	  |   三角矩阵   |
|   	tp   | triangular matrix (packed storage)  | 三角矩阵（压缩存储)     |
|   	tb	   | triangular band matrix  |    三角带状矩阵  |

- lda，ldb，ldc：矩阵的前导维数 (主维数)
- incx，incy：x,y的增量或步长, 一般设置为1
- X* = conjg(X) ：求X的共轭
- Order : 指定行-主©或列-主(Fortran)数据排序。
- Uplo : 指定是使用矩阵中的上三角还是下三角。有效值是“U”或“L”。
- TransA : 指定是使用矩阵A (‘N’或’n’)还是A (‘T’、‘t’)的转置，共轭（‘C’或’c’）。
- Diag : 指定矩阵是否是单位三角形。可能的值是“U”(单位三角形)或“N”(非单位三角形)。
- Side: 决定矩阵相乘的顺序，Side=L（AB） Side=R（BA）。
- Hermitian 埃尔米特矩阵（又称“自共轭矩阵”）是共轭对称的方阵。埃尔米特矩阵中每一个第i 行第j 列的元素都与第j 行第i 列的元素的共轭相等。n阶复方阵A的对称单元互为共轭，即A的共轭转置矩阵等于它本身，则A是埃尔米特矩阵(Hermitian Matrix)。显然埃尔米特矩阵是实对称矩阵的推广
