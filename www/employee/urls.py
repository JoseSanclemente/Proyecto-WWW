from rest_framework import routers

from .viewsets import EmployeeViewSets

router = routers.SimpleRouter()
router.register('employee', EmployeeViewSets)

urlpatterns = router.urls
