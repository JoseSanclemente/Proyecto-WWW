from django.db import models
import uuid

# Create your models here.
class Reads(models.Model):
    reads_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    meter = models.ForeignKey('meter.Meter', null=True, on_delete=models.SET_NULL)
    active_value = models.BooleanField()
    reactive_value = models.BooleanField()
    date = models.DateField()

