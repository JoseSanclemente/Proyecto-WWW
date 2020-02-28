PROJECTS = user invoice period subestation transmodel contract client
makemigrations:
	$(foreach project, $(PROJECTS), python manage.py makemigrations $(project);)