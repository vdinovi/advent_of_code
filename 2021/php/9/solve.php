#!/usr/bin/env php
<?php
$grid = array();
while ($line = fgets(STDIN)) {
    $chars = str_split(rtrim($line));
    $row = array();
    foreach ($chars as $char) {
        array_push($row, (int)$char);
    }
    array_push($grid, $row);
}
$height = sizeof($grid);
$width = sizeof($grid[0]);

$visited = array();
$product_of_basin_sizes = 1;
$basin_sizes = array();
for ($y = 0; $y < $height; $y++) {
    for ($x = 0; $x < $width; $x++) {
        $h = $grid[$y][$x];
        if ($y > 0           && $grid[$y - 1][$x] <= $h) { continue; }
        if ($x < $width - 1  && $grid[$y][$x + 1] <= $h) { continue; }
        if ($y < $height - 1 && $grid[$y + 1][$x] <= $h) { continue; }
        if ($x > 0           && $grid[$y][$x - 1] <= $h) { continue; }

        $size = 0;
        $basin_points = array(array($y, $x));
        while (sizeof($basin_points) > 0) {
            $point = array_shift($basin_points);
            $y1 = $point[0];
            $x1 = $point[1];
            if (isset($visited[sprintf("%d,%d", $y1, $x1)])) { continue; }

            $size += 1;
            $visited[sprintf("%d,%d", $y1, $x1)] = true;
            if ($y1 > 0 && $grid[$y1 - 1][$x1] > $h && $grid[$y1 - 1][$x1] < 9) { array_push($basin_points, array($y1 - 1, $x1)); }
            if ($x1 < $width - 1  && $grid[$y1][$x1 + 1] > $h && $grid[$y1][$x1 + 1] < 9) { array_push($basin_points, array($y1, $x1 + 1)); }
            if ($y1 < $height - 1 && $grid[$y1 + 1][$x1] > $h && $grid[$y1 + 1][$x1] < 9) { array_push($basin_points, array($y1 + 1, $x1));  }
            if ($x1 > 0 && $grid[$y1][$x1 - 1] > $h && $grid[$y1][$x1 - 1] < 9) { array_push($basin_points, array($y1, $x1 - 1)); }
        }
        array_push($basin_sizes, $size);
    }
}

rsort($basin_sizes);
$product_of_largest_basin_sizes = 1;
for ($i = 0; $i < 3; $i++) {
    $product_of_largest_basin_sizes *= $basin_sizes[$i];

}
printf("%d\n", $product_of_largest_basin_sizes);
?>
