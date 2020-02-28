import uuid
from django.db import models

# Create your models here.


class Period(models.Model):
    period_id = models.UUIDField(primary_key=True, unique=True, default=uuid.uuid4, editable=False)
    start_date = models.DateField()
    end_date = models.DateField()
    kw_h_price = models.FloatField()
    iva = models.FloatField()

