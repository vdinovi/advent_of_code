#!/usr/bin/perl

$x = 0;
$y = 0;

foreach my $line (<STDIN>) {
    chomp $line;
    if ($line =~ /\s*(\w+) (\d+)\s*/) {
        my ($dir, $dist) = ($1, int($2));
        if ($dir eq "forward") { $x += $dist; }
        elsif ($dir eq "down") { $y += $dist; }
        elsif ($dir eq "up")   { $y -= $dist; }
    }
}

print $x * $y, "\n"
