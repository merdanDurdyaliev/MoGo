# MoGo

## Switch from `main` to `master`

bash
git branch -m main master        # rename branch ---
git push -u origin master        # push new branch
git push origin --delete main    # delete old main (optional)
Now your default branch is master.

## Run the website locally
1. Make sure dependencies are installed.
go mod tidy

2. Run your server:
go run main.go

3. Open in browser:
cpp
http://127.0.0.1:4001
Save changes
bash
git add .
git commit -m "Your message"
git push origin master
