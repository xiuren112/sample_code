# =============================================
# Sample R Code
# Author: https://github.com/xiuren112
# Date: 2026-01-05
# Description: Utility functions with basic error handling and examples
# =============================================

# -----------------------------
# Function 1: Add two numbers
# -----------------------------
add_numbers <- function(a, b) {
  if (!is.numeric(a) || !is.numeric(b)) {
    stop("Both 'a' and 'b' must be numeric")
  }
  return(a + b)
}

# Example usage:
# result <- add_numbers(2, 3)
# print(result)  # 5

# -----------------------------
# Function 2: Count words in a string
# -----------------------------
count_words <- function(text) {
  if (!is.character(text)) {
    stop("'text' must be a character vector")
  }
  # Count non-space sequences
  return(sum(gregexpr("\\S+", text)[[1]] > 0))
}

# Example usage:
# words <- count_words("Hello world from R")
# print(words)  # 4

# -----------------------------
# Function 3: Safe division
# -----------------------------
safe_divide <- function(a, b) {
  if (!is.numeric(a) || !is.numeric(b)) {
    stop("Both 'a' and 'b' must be numeric")
  }
  if (b == 0) {
    warning("Division by zero! Returning NA.")
    return(NA)
  }
  return(a / b)
}

# Example usage:
# print(safe_divide(10, 2))  # 5
# print(safe_divide(10, 0))  # NA with warning
