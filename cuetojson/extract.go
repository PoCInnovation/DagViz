package cuetojson

func ExtractInfos(programs []CueProgram) []CueInfos {
	var infos []CueInfos

	for _, program := range programs {
		info := CueInfos{}
		info.Values = program.Values

		infos = append(infos, info)
	}
	return infos
}
