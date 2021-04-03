class Solution {
private:
    set<string> dict;
    map<string, int> dis;
public:
    void init() {
        dict.clear();
        dis.clear();
    }
    void bfs(string start, string end ) {
        queue<string> q;
        q.push(start);
        while(!q.empty()) {
            string x = q.front();
            q.pop();
            int d = dis[x];
            for(int i = 0; i < x.size(); i ++) {
                char tmp = x[i];
                for (char j = 'a'; j <= 'z'; ++j) {
                    if (x[i] == j) {
                        continue;
                    }
                    x[i] = j;
                    if (dict.find(x) == dict.end()) {
                        continue;
                    }
                    if (dis.find(x) == dis.end()) {
                        dis[x] = d+1;
                        q.push(x);
                    }
                }
                x[i] = tmp;
            }
        }
    }
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        init();
        dict.insert(beginWord);
        for (auto &s : wordList) {
            dict.insert(s);
        }
        dis[beginWord] = 0;
        bfs(beginWord, endWord);
        if (dis.find(endWord) == dis.end()) {
            return 0;
        }
        return dis[endWord]+1;
    }
};