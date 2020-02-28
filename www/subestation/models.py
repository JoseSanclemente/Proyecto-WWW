from django.db import models
import uuid

# Create your models here.
class Subestation(models.Model):
    sub_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    name = models.CharField(max_length=30)
    address = models.CharField(max_length=30)
    active = models.BooleanField()
    