#!/usr/bin/env perl
use strict;
use warnings;

# Devel::Cover's command-line logic is contained in Devel::Cover::Web or its core.
# Loading the module and running its standard CLI entry point:
use Devel::Cover;

# Alternate approach: If you want to explicitly run the contents of the upstream script,
# you can require it from the library path:
require 'bin/cover';
