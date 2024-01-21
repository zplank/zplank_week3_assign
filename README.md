# zplank_week3_assign

When comparing the output of the R/Python linear regression models and to the results in Go using the Anscombe Quartet sample data, we see similar outcomes. Aside from some rounding, the below tables show that Go produces the same performance results of each of the 4 models. 

Linear Regression Model Output Stats: 

--Intercept 

	    | x1~y1	| X2~y2  | x3~y3  | x4~y4
------------------------------------------
Python 	| 3.001 | 3.0009 | 3.0025 | 3.0017
------------------------------------------		
R		| 3.001 | 3.001  | 3.0025 | 3.0017
------------------------------------------		
Go 		| 3.001 | 3.0009 | 3.0025 |	3.0017	
				
--Slope 

	    | x1~y1	 | X2~y2  | x3~y3  | x4~y4
------------------------------------------
Python 	| 0.5001 | 0.500  | 0.4997 | 0.4999
------------------------------------------		
R		| 0.5001 | 0.500  | 0.4997 | 0.4999
------------------------------------------		
Go 		| 0.5001 | 0.5000 | 0.4997 | 0.4999	

--R^2 

	    | x1~y1	 | X2~y2  | x3~y3  | x4~y4
------------------------------------------
Python 	| 0.6667 | 0.666  | 0.666  | 0.667
------------------------------------------		
R		| 0.6665 | 0.6662 | 0.6663 | 0.6667
------------------------------------------		
Go 		| 0.6665 | 0.6662 | 0.6663 | 0.6667	


When it comes to execution times between Go, Python, and R, Python and R are almost instatneous whereas Go is not. Go takes around 7-10 seconds per model and similarly, the overall model took approximately 7.5 seconds to execute. When comparing the three programs, Go is the slowest to execute basic queries. 


Overall, my recommendation would be to begin utilizing Go across the company. These tests determined that simple linear regression models is possible using Go perform similarly between all languages. One thing to take into consideration with utilizing Go to replace Python or R is that these tests were only performed on simple linear regression models and not more complex models. Go packages could differ in how models perform when multiple variables are introduced. In addition, execution times lagged with Go when testing on these simple models. That being said, 7-10 seconds is not a dealbreaker when it comes to executing queries and these times could differ across all three programs when testing on different types of models. If the endgoal is to push everyone to utilize the same programing language, Go has proven that it is capable to handle linear regression models, producing the same output as Python and R.  