#!/usr/bin/env awk -f
# coverage.awk 脚本用于检查代码覆盖率是否达到阈值
# 使用方法: go tool cover -func=coverage.out | awk -v target=1 -f coverage.awk

BEGIN {
    total_coverage = 0
}

/total:/ {
    # 解析 total: 行，格式: total:                                                                   (statements)                                     2.0%
    # 提取百分比数值
    if (match($0, /[0-9.]+%/)) {
        # 提取匹配的字符串
        match_str = substr($0, RSTART, RLENGTH)
        # 移除 % 符号并转换为数字
        gsub(/%/, "", match_str)
        total_coverage = match_str
    }
}

END {
    if (total_coverage == 0) {
        print "ERROR: Could not determine total coverage"
        exit 1
    }
    
    if (total_coverage < target) {
        printf "FAIL: test coverage is %.1f%% (quality gate is %.1f%%)\n", total_coverage, target
        exit 1
    } else {
        printf "PASS: test coverage is %.1f%% (quality gate is %.1f%%)\n", total_coverage, target
        exit 0
    }
}

