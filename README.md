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

--insert stats about timing 

Overall, my recommendation would be to begin utilizing Go across the company. These tests determined that simple linear regression models is possible using Go perform similarly between all languages. One thing to take into consideration with utilizing Go to replace Python or R is that these tests were only performed on simple linear regression models and not more complex models. Go packages could differ in how models perform when multiple variables are introduced. 