# How to run
# Full teardown (safe, preserves PVCs/data by default):

FULL_RESET=true ./clean-event-streaming-prod.sh





Delete just topics (keep cluster):

ONLY_TOPICS=true ./clean-event-streaming-prod.sh




Full teardown + wipe data volumes (irreversible – removes all topic data):

FULL_RESET=true WIPE_PVC=true ./clean-event-streaming-prod.sh


If you ever change the cluster name, you can also force it:

CLUSTER_NAME=small-kafka FULL_RESET=true ./clean-event-streaming-prod.sh
