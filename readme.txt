heapvis - visualizing heap profile diffs over time


THE HYPOTHESIS

	while being able to visualize a given profile taken at a certain point
	in time is great, the great insights will come from looking at how that
	changes over time, as time passes.

	`heapvis` aims at making that simple to do.



IMPLEMENTATION

	vis

		- sort by the biggest difference over time
		- delta + variance?



	- frequency trails: 

	  	https://flowingdata.com/charttype/frequency-trails/

		- with matplotlib, it seems like it can be done quite nicely
			https://stackoverflow.com/questions/17614499/frequency-trail-in-matplotlib



	- seems like it could be adequatable to go with an annotated heatmap too
		http://www.brendangregg.com/FrequencyTrails/intro.html
