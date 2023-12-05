
function optimize(positions, fn)
    solution = NaN
    error = Inf

    for soln in 0 : maximum(positions) - 1 
        err = reduce(+, [fn(pos, soln) for pos in positions], init=0)
        if err < error
            error = err
            solution = soln
        end
        #println("solution $soln has error $err")
    end
    return solution, error
end

function abs_diff(x, y)
    return abs(x - y)
end

function linear_abs_diff(x, y)
    return sum([abs(diff) for diff = 0 : abs(x-y)])
end

positions = map(n -> parse(Int64, n), split(readline(), ","))
solution, error = optimize(positions, linear_abs_diff)
println("Optimal solution is $solution with error $error")
