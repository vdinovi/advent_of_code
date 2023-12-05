#!/usr/bin/perl

$x = 0;
$y = 0;
$aim = 0;

sub Advance {
    my ($dir, $dist) = @_;
    if ($dir eq "forward") { $x += $dist; $y += $aim * $dist }
    elsif ($dir eq "down") { $aim += $dist; }
    elsif ($dir eq "up")   { $aim -= $dist; }
}

foreach my $line (<STDIN>) {
    chomp $line;
    if ($line =~ /\s*(\w+) (\d+)\s*/) {
        my ($dir, $dist) = ($1, int($2));
        Advance($dir, $dist);
    }
}

print $x * $y, "\n"
