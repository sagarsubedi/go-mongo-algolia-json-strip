# MongoDB to Algolia json transformer

This little script takes in a json file that contains data extracted from mongodb.

It then strips unnecessary fields and generates a new json file with the necessary fields.

You know what I learned from this?

- Not how to parse json (already knew it)
- Not how to generate new json with only necessary data

I actually had this exact same script in JS. Since I was learning Go at the time, I thought of changing it to Go code and see how fast it works.

To my surprise, the go code I had actually took a few milliseconds more than the JS. I couldn't believe myself so I ran the test about 20 minutes. On average, Node did it in about `10ms` while Go took about `13ms`. I couldn't believe this and started looking into the parser I was using (since that was the obvious culprit because that was the only thing I was doing).

After some research, I realized that i was wasting a lot of memory by loading it all into memory, then creating a variety of slices, etc. Turns out there is another package called `json-iterator/go`. I used it as a drop-replacement only. And it reduced the average time from `13ms` to `4ms`. Finally, I liked what I saw and I stopped it there. Give `jsoniter` a try. It has almost the same API as the built-in "encoding/json" package, but is faster. Anywas, Go still remains my favorite.
