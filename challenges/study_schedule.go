package challenges

func Study_schedule(permanence_period [][]int, target_time int) int {
	students_online := 0

	for i := 0; i < len(permanence_period); i++ {
		if permanence_period[i][0] <= target_time && target_time <= permanence_period[i][1] {
			students_online += 1
		}
	}

	return students_online
}
