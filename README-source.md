# dev-to-github-actions-readme

# 🚀 Update your dev.to articles into your Github profile using Go Templates, Docker, GitHub Actions 🚀

https://github.com/coding-to-music/dev-to-github-actions-readme

From / By https://github.com/huantt/article-listing

https://dev.to/jacktt/update-your-devto-articles-into-your-github-profile-4dpi

<!-- <div style="text-align:center;">
  <img src="/images/chakra.jpg" alt="Image" />
  <p><em>Chakra Component Library with Next.js</em></p>
</div> -->

## Go Environment:

```java
Go Version
```

Output

```java
go version go1.17.6 linux/amd64
```

## Environment variables:

```java
none
```

## GitHub

```java
git init
git add .
git remote remove origin
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:coding-to-music/dev-to-github-actions-readme.git
git push -u origin main
```

## Install latest Go version

Latest downloads are available here:

https://go.dev/dl/

```java
wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
```

### Extract the Archive:

Use the tar command to extract the contents of the downloaded archive. Replace "your-username" with your actual username:

```java
sudo tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
```

This command will extract the Go files into the /usr/local directory.

### Set Go Environment Variables:

You need to set Go's environment variables so that your system knows where to find the Go binaries. Add the following lines to your shell profile file (e.g., ~/.bashrc, ~/.zshrc, or ~/.profile), replacing "your-username" with your actual username if necessary:

```java
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### After adding these lines, run the following command to apply the changes immediately to your current session:

```java
source ~/.bashrc  # or ~/.zshrc, or the appropriate file for your shell
```

### Verify the Installation:

After setting the environment variables, you can verify the installation by running:

```java
go version
```

It should display the Go version you downloaded and installed.

That's it! You've successfully installed Go from the tar.gz archive on your Linux system. You can now start using Go for your development projects.
