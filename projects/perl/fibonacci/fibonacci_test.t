use strict;
use warnings;

use Test::More tests => 13;

use projects::perl::fibonacci::fibonacci;

my @numbers = (0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144);

for my $i (0 .. $#numbers) {
  my $result = fibonacci::fibonacci($i);
  my $expected = $numbers[$i];
  is($expected, $result, "fibonacci($i) == $expected");
}
