# Create git commit and push of current repo with username and password
# Add Usename and passward
git:
	git add .
	git commit -m "$m"
	git push -u origin main