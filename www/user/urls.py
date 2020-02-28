from rest_framework import routers

from .viewsets import UserViewSets

router = routers.SimpleRouter()
router.register('user', UserViewSets)

urlpatterns = router.urls
