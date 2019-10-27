heapvis - visualizing heap profile diffs over time


THE HYPOTHESIS

	while being able to visualize a given profile taken at a certain point
	in time is great, the great insights will come from looking at how that
	changes over time, as time passes.

	`heapvis` aims at making that simple to do.



IMPLEMENTATION

	
	vis


		y = function name (sorted by biggest diff in delta)
		x = delta over time



		csv:

			main.fn,main.main,fmt.Printf
			1,2,3
			4,2,1
			3.2.1



	- frequency trails: 

	  	https://flowingdata.com/charttype/frequency-trails/

		- with matplotlib, it seems like it can be done quite nicely
			https://stackoverflow.com/questions/17614499/frequency-trail-in-matplotlib


		- ridgeplots: https://www.data-to-viz.com/graph/ridgeline.html



	- seems like it could be adequatable to go with an annotated heatmap too
		http://www.brendangregg.com/FrequencyTrails/intro.html
