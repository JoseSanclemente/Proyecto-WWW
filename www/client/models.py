import uuid
from django.db import models

# Create your models here.


class Client(models.Model):
    client_id = models.UUIDField(primary_key=True, unique=True, default=uuid.uuid4, editable=False)
    document_id = models.CharField(max_length=30)
    email = models.CharField(max_length=30)
    name = models.CharField(max_length=30)
    phone_number = models.CharField(max_length=15)
    address = models.CharField(max_length=30)
    active = models.BooleanField()
