from django.db import models
import uuid

# Create your models here.
class User(models.Model):
    user_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    national_id = models.CharField(max_length=19)
    name = models.CharField(max_length=30)
    date_of_birth = models.DateField()
    email = models.CharField(max_length=30)
    address = models.CharField(max_length=30)
    active = models.BooleanField()
    role = models.CharField(max_length=15)
