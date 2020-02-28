from rest_framework import routers

from .viewsets import SubestationViewSets

router = routers.SimpleRouter()
router.register('Subestation', SubestationViewSets)

urlpatterns = router.urls
