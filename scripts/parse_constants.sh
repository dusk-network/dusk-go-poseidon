#!/bin/bash

SCRIPT=`realpath -s $0`
SCRIPT_PATH=`dirname $SCRIPT`
PROJECT_ROOT=`dirname $SCRIPT_PATH`
INTERNAL_PATH="$PROJECT_ROOT/internal"
ARK_PATH="$INTERNAL_PATH/ark.bin"
MDS_PATH="$INTERNAL_PATH/mds.bin"

CONSTANTS_LEN=`xxd -c 32 -g 4 $ARK_PATH | wc -l`
MDS_SIZE=`ls -l $MDS_PATH | awk '{ print $5 }'`
export MDS_WIDTH=`echo "sqrt($MDS_SIZE / 32)" | bc`

echo "package poseidon"
echo ""
echo "import ("
echo "	ristretto \"github.com/bwesterb/go-ristretto\""
echo ")"
echo ""
echo "// DefaultRoundKeys extracted from Rust Merkle Tree implementation"
echo "var DefaultRoundKeys = [$CONSTANTS_LEN]ristretto.Scalar{"

xxd -c 32 -g 1 $ARK_PATH | awk '{ print "ristretto.Scalar { 0x" $5 $4 $3 $2 ", 0x" $9 $8 $7 $6 ", 0x" $13 $12 $11 $10 ", 0x" $17 $16 $15 $14 ", 0x" $21 $20 $19 $18 ", 0x" $25 $24 $23 $22 ", 0x" $29 $28 $27 $26 ", 0x" $33 $32 $31 $30 " }," }'

echo "}"
echo ""
echo "// DefaultMDSMatrix extracted from Rust Merkle Tree implementation"
echo "var DefaultMDSMatrix = [][]ristretto.Scalar{"

xxd -c 32 -g 1 $MDS_PATH | awk '{ print "ristretto.Scalar { 0x" $5 $4 $3 $2 ", 0x" $9 $8 $7 $6 ", 0x" $13 $12 $11 $10 ", 0x" $17 $16 $15 $14 ", 0x" $21 $20 $19 $18 ", 0x" $25 $24 $23 $22 ", 0x" $29 $28 $27 $26 ", 0x" $33 $32 $31 $30 " }," }' | perl -e '$w = $ENV{'MDS_WIDTH'}; $c = 1; while(<>) { if($c == 1) { print "[]ristretto.Scalar{\n"; } print; if($c == $w){print"},\n"; $c = 0;} $c++; }'

echo "}"
