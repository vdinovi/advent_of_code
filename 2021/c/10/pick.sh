#!/usr/bin/env bash

sort -n | awk '
    BEGIN {
        count = 0;
    }

    {
        costs[count++] = $1;
    }

    END {
        if ((count % 2) == 1) {
          median = costs[int(count / 2)];
        } else {
          median = (costs[count / 2] + costs[count / 2 - 1]) / 2;
        }
        print median;
    }
'

