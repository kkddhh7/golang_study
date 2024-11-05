import numpy as np

# 행렬 A 정의
A = np.array([[1, -2, 3, 5],
              [2, 2, -1, 0],
              [3, 0, 1, 2],
              [1, 0, 2, 0]])

# 2A 계산
two_A = 2 * A
print("2A:\n", two_A)

# A^T 계산
A_transpose = np.transpose(A)
print("\nA^T:\n", A_transpose)

# A^-1 계산 (만약 가능한 경우)
try:
    A_inverse = np.linalg.inv(A)
    print("\nA^-1:\n", A_inverse)
except np.linalg.LinAlgError:
    print("\nA는 역행렬이 존재하지 않습니다.")

# 선형독립성 확인을 위한 행렬식 계산
det_A = np.linalg.det(A)
if det_A != 0:
    print("\nA는 선형독립입니다 (det(A) = {}).".format(det_A))
else:
    print("\nA는 선형독립이 아닙니다 (det(A) = 0).")