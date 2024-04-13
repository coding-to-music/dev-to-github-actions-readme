## About
Collect your latest articles from sources such as [dev.to](https://dev.to), and then update the `README.md`.

## Use GitHub Action to update your README

**Step 1:** In your repository, create a file named `README.md.template`.

**Step 2:** Write anything you want within the `README.md.template` file.

**Step 3:** Embed one of the following entities within your `README.md.template`:

- **Article listing:**
```shell
{{ template "article-list" .Articles }}
```
- **Article table:**
```shell
{{ template "article-table" .Articles }}
```

If you are familiar with Go templates, you have access to the `root` variable, which includes the following fields:

- `Articles`: An array of Article. You can view the Article struct definition in [model/article.go](model/article.go).
- `Time`: Updated Time
- `Author`: Author of articles

**Step 4**: Register Github Action
- Create a file `.github/workflows/update-articles.yml` in your repository.
```yml
name: "Cronjob"
on:
schedule:
- cron: '15 0 * * *'

jobs:
    update-articles:
        permissions: write-all
        runs-on: ubuntu-latest
        steps:
            - uses: actions/checkout@v3
            - name: Generate README
              uses: coding-to-music/dev-to-github-actions-readme@v1.0.0
              with:
                username: YOUR_USERNAME_ON_DEV_TO                
                template-file: 'README.md.template'
                out-file: 'README.md'
                limit: 5
            - name: Commit
              run: |
                if git diff --exit-code; then
                    echo "No changes to commit."
                    exit 0
                else
                    git config user.name github-actions
                    git config user.email github-actions@github.com
                    git add .
                    git commit -m "update"
                    git push origin main
                fi
```

**Step 5**: Commit your change, then Github actions will run as your specified cron to update Articles into your README.md file

## Below is my recent articles Tom Connors collected from dev.to
### Table


<table>
        <tr>
            <td width="300px">
                <a href="https://dev.to/codingtomusic/mits-premier-programming-competition-battlecode-open-to-all-registration-is-now-open-4f0"><img src="https://media.dev.to/cdn-cgi/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2F6f85sgi5wpwc66ftrtur.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/codingtomusic/mits-premier-programming-competition-battlecode-open-to-all-registration-is-now-open-4f0">MIT&#39;s Premier Programming Competition, Battlecode, open to all, registration is...</a>
                <div>It&#39;s Battlecode Season, open to all   Battlecode 2023 begins on January 9th, 2023!  Register...</div>
                <div><i>05/12/2022</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px">
                <a href="https://dev.to/codingtomusic/playing-around-with-openais-chatgpt-3-to-ask-how-to-upload-multiple-images-to-aws-s3-with-mongodb-and-firebase-auth-48h0"><img src="https://media.dev.to/cdn-cgi/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fu0bmon6vlubukgk59hkb.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/codingtomusic/playing-around-with-openais-chatgpt-3-to-ask-how-to-upload-multiple-images-to-aws-s3-with-mongodb-and-firebase-auth-48h0">Playing around with Open.ai&#39;s ChatGPT-3 to ask how to upload...</a>
                <div>It Works!   https://chat.openai.com/chat  Question 1    use typescript to put multiple...</div>
                <div><i>05/12/2022</i></div>
            </td>
        </tr>
        <tr>
            <td width="300px">
                <a href="https://dev.to/codingtomusic/battlecode-2021-overview-and-instructions-create-java-bots-in-a-virtual-world-4ibn"><img src="https://media.dev.to/cdn-cgi/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fi%2Fee9guqjz5g4qaoqcg882.png" alt="thumbnail"></a>
            </td>
            <td>
                <a href="https://dev.to/codingtomusic/battlecode-2021-overview-and-instructions-create-java-bots-in-a-virtual-world-4ibn">BattleCode 2021 - Overview and instructions - create java bots...</a>
                <div>BattleCode 2021 - Overview                     BattleCode 2021 - Overview   Introduction Why...</div>
                <div><i>15/01/2021</i></div>
            </td>
        </tr>
</table>


### List

- [MIT&#39;s Premier Programming Competition, Battlecode, open to all, registration is...](https://dev.to/codingtomusic/mits-premier-programming-competition-battlecode-open-to-all-registration-is-now-open-4f0) - 05/12/2022
- [Playing around with Open.ai&#39;s ChatGPT-3 to ask how to upload...](https://dev.to/codingtomusic/playing-around-with-openais-chatgpt-3-to-ask-how-to-upload-multiple-images-to-aws-s3-with-mongodb-and-firebase-auth-48h0) - 05/12/2022
- [BattleCode 2021 - Overview and instructions - create java bots...](https://dev.to/codingtomusic/battlecode-2021-overview-and-instructions-create-java-bots-in-a-virtual-world-4ibn) - 15/01/2021

*Updated at: 2024-04-13T12:21:17Z*