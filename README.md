# time2duration

Write a Go cli that reads from the stdin stream line by line, and print to the stdout "${duration}\t${line}\n" for each line.

https://chat.openai.com/share/00f2ee1b-df0f-47af-93b9-33fd6de958d4

Notes

1.  > Me: ... before the program exits, can it print the top 5 longest duration in desc order
    >     "Top 5 longest duration:\n${duration}\t${line}\n${duration}\t${line}\n${duration}\t${line}\n${duration}\t${line}\n${duration}\t${line}\n"
    
    after several to-and-fro later, we spotted a problem
    
    > Me: is `durationLines` growing unbounded? shouldn't we always only just keep the longest 5 durations?
    
    😮‍💨

2. As we iterate on the code, notice that we're not always making forward progress. e.g. at one point, the code regressed and had to be manually caught

    > Me: we seem to have forgotten to print lines while we read [from stdin]?
    > 
    > ChatGPT: Indeed, it seems that in focusing on maintaining the top 5 durations in the heap, the part of the code that prints each line along with its corresponding duration was omitted. Let's reintegrate that functionality back into the code so that it prints each line and its duration immediately when processed, as well as maintaining and finally printing the top 5 longest durations at the end.

    There's always this tension of _"I should keep whatever I obtained and salvage by writing the rest myself instead of pushing ChatGPT further risking worse bugs."_
