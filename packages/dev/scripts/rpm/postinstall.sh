#!/bin/bash

# Copyright 2023 Adam Chalkley
#
# https://github.com/atc0005/send2teams
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

project_org="atc0005"
project_shortname="send2teams"

project_fq_name="${project_org}/${project_shortname}"
project_url_base="https://github.com/${project_org}"
project_repo="${project_url_base}/${project_shortname}"
project_releases="${project_repo}/releases"
project_issues="${project_repo}/issues"
project_discussions="${project_repo}/discussions"



echo
echo "Thank you for installing packages provided by the ${project_fq_name} project!"
echo
echo "#######################################################################"
echo "NOTE:"
echo
echo "This is a dev build; binaries installed by this package have a _dev"
echo "suffix to allow installation alongside stable versions."
echo
echo "Feedback for all releases is welcome, but especially so for dev builds."
echo "Thank you in advance!"
echo "#######################################################################"
echo
echo "Project resources:"
echo
echo "- Obtain latest release: ${project_releases}"
echo "- View/Ask questions: ${project_discussions}"
echo "- View/Open issues: ${project_issues}"
echo
