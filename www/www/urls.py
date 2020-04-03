"""www URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/3.0/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('user.urls')),
    path('', include('employee.urls')),
    path('', include('subestation.urls')),
    path('', include('transformer.urls')),
    path('', include('transmodel.urls')),
    path('', include('client.urls')),
    path('', include('client.urls')),
    path('', include('contract.urls')),
    path('', include('period.urls')),
    path('', include('meter.urls')),
    path('', include('metmodel.urls')),
    path('', include('reads.urls')),
    path('', include('payment.urls')),
    path('', include('invoice.urls')),
]
