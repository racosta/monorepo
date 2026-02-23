use v5.40;

package fibonacci;

sub fibonacci {
  my ($n) = @_;

  if ( $n < 2 ) {
    return $n;
  }

  my $n1 = 0;
  my $n2 = 1;

  for my $i ( 2 .. $n ) {
    my $sum = $n1 + $n2;
    $n1 = $n2;
    $n2 = $sum;
  }

  return $n2;
}

1;
