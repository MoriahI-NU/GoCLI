#!/Program Files/R/R-4.3.1/bin/x64/Rgui Rscript

#number of runs
N = 100 

# start keeping track of time
start_time <- Sys.time()

# create and access output file
sink("housesOutputR.txt")

#read input csv and print summary stats to the output file - repeated 100 times
for (i in 1:N) {
    houses = read.csv(file = "housesInput.csv", header = TRUE)
    print(summary(houses)) 
}

#leave the output file
sink()

#calculate execution time for all runs
end_time <- Sys.time()
runtime <- end_time - start_time

cat("Execution Time (sec) for 100 runs = ", runtime)
cat("\nAverage Execution Time (sec) per run = ", runtime/100)