from django.db import models

# Create your models here.


class User(models.Model):
    national_id = models.CharField(max_length=19)
    name = models.CharField(max_length=30)
    date_of_birth = models.DateField()
    email = models.CharField(max_length=30)
    address = models.CharField(max_length=30)
    active = models.BooleanField()
    role = models.CharField(max_length=15)
