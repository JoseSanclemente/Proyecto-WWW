from rest_framework import routers

from .viewsets import PeriodViewSets

router = routers.SimpleRouter()
router.register('contract', PeriodViewSets)

urlpatterns = router.urls
