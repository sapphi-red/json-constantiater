reset
set datafile separator ","
set style data histogram
set style histogram cluster gap 1
set style fill solid 0.7 border
set border lw 0.8
set boxwidth 0.8

set xtics nomirror rotate by 90 scale 0 right
set ytics 200 nomirror rotate by 90
set yrange [0:]

set xlabel "Medium" rotate by 180
set ylabel "ns / op" offset 1

unset key

set grid

set term pngcairo font "Times Roman,14"  enhanced size 400,800
set output "./medium.png"

plot 'medium.csv' using 2:xticlabels(1) title columnheader(2) lc rgb "#23cc42"
replot
