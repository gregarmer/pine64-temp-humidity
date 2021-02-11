set title 'Atmospheric Conditions' font ',18'
set term pngcairo enhanced font 'Roboto Condensed' size 1920,1024
set output 'pmsd001.png'
set grid xtics ytics
set timestamp '%Y-%m-%d %T%z' font ',8'

set xdata time
set timefmt '%s'
set xlabel 'time [HH:MM]'
set xtics format '%H:%M' time
set mxtics 6

set link x2
set x2label 'date [yyyy-mm-dd]'
set x2tics nomirror 1609459200,86400 format '%Y-%m-%d' time

set ylabel 'temperature [°C]'
set yrange [-10:40]
set ytics nomirror -20,5
set mytics 5

set y2label 'relative humidity [%]' rotate by -90
set y2range [20:70]
set y2tics nomirror 0,5
set my2tics 5

plot \
0 w l lc rgb 'black' t 'freezing point', \
'/var/log/pmsd001.log' u 1:2 w l lw 2 lc rgb '#d50000' t 'temperature', \
'' u 1:4 w l lw 2 lc rgb '#0091ea' axes x1y2 t 'relative humidity'
