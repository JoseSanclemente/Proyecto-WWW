from rest_framework import routers

from .viewsets import TransmodelViewSets

router = routers.SimpleRouter()
router.register('Transmodel', TransmodelViewSets)

urlpatterns = router.urls
