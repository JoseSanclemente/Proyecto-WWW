# Generated by Django 3.0.3 on 2020-04-03 00:28

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('client', '0002_auto_20200403_0009'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='client',
            name='date_of_birth',
        ),
    ]
