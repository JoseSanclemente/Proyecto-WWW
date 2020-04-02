""" from rest_framework import routers

from .viewsets import UserViewSets

router = routers.SimpleRouter()
router.register('user', UserViewSets)

urlpatterns = router.urls """


from django.urls import path
from user import views

urlpatterns = [
    path('login/', views.UserAPIView.as_view()),
    path('captcha/', views.CaptchaAPIView.as_view()),
]