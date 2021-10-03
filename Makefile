# Create git commit and push of current repo with username and password
# makefile
git:
	git add .
	git commit -m "$m"
	git push -u origin main