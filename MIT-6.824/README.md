# MIT 6.824 Distributed Systems

Graduate Course on Distributed Systems by MIT

Course Home: http://nil.csail.mit.edu/6.824/2017/
Lectures: https://www.youtube.com/playlist?list=PLpl804R-ZwjKCOwWpTZ21eeaBS3uBrMfV (sound is aweful)
Past Exams: http://nil.csail.mit.edu/6.824/2017/quizzes.html

## Week 1: Map Reduce
Paper
http://nil.csail.mit.edu/6.824/2017/papers/mapreduce.pdf

My Notes:
- It uses a restricted programming model involving defining a job as a map step and a reduce step.
- The programming model is really easy and allows people without experience in distributed or parallel systems to reap the benefits of a parallel distributed system.
- Flow looks like: Input Files -> Divided amongst Mappers -> Input File -> Reader -> Mapper.Map(<Key1, Value1>) -> Emits <Key2, Value2> -> Paritioned -> Reducer.Reduce(<Key2, List<Value2>>) -> Output File per Reducer
- It is written as a library that a client program interacts with.
- The job is configured to run on M + R + 1 workers, with M mappers, R reducers and a master.
- The master deals with allocating the jobs, handling failure, workers that are straggling.
- It deals with a large amount of input and output data.
- It uses a large ammount of commidity hardware, commodity machines with commodity disks and network.
- There are 2x failure recovery mechanisms:
  1) Failed jobs are re-ran on different workers in the cluster.
  2) Records that fail often are marked and skipped.
- If a worker is straggling on a job the master can schedule a backup execution, which ever finishes first wins.
- Not all problems are able to be expressed in terms of a map / reduce job. But a lot of useful problems can be.

(Not map reduce, but interesting, reference [4] in paper: Google Search Architecture: https://static.googleusercontent.com/media/research.google.com/en//archive/googlecluster-ieee.pdf)

## Lab 1: Map Reduce
http://nil.csail.mit.edu/6.824/2017/labs/lab-1.html

## Related Video:
Cluster Computing and MapReduce Lecture 1
https://www.youtube.com/watch?v=yjPBkvYh-ss&t=1s
- Moores law plateau, but even if CPUs keep getting faster, memory bandwidth seems to be the bottleneck
- Implies a limit on what will be able to be achieved by a single computer
